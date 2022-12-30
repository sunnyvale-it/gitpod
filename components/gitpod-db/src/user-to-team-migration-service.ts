/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { Team, User } from "@gitpod/gitpod-protocol";
import { ErrorCodes } from "@gitpod/gitpod-protocol/lib/messaging/error";
import { inject, injectable } from "inversify";
import { ProjectDB } from "./project-db";
import { TeamDB } from "./team-db";
import { ResponseError } from "vscode-jsonrpc";
import { WorkspaceDB } from "./workspace-db";
import { TypeORM } from "./typeorm/typeorm";
import { log, LogContext } from "@gitpod/gitpod-protocol/lib/util/logging";

@injectable()
export class UserToTeamMigrationService {
    @inject(TeamDB) protected readonly teamDB: TeamDB;
    @inject(ProjectDB) protected readonly projectDB: ProjectDB;
    @inject(WorkspaceDB) protected readonly workspaceDB: WorkspaceDB;
    @inject(TypeORM) protected readonly typeorm: TypeORM;

    async migrateUser(user: User): Promise<void> {
        if (!(await this.needsMigration(user))) {
            return;
        }
        await this.createTeamOfOne(user);
    }

    async createTeamOfOne(user: User): Promise<Team> {
        const ctx: LogContext = { userId: user.id };
        log.info(ctx, "Creating team of one.");
        let team;
        let tries = 0;
        while (!team && tries++ < 10) {
            try {
                let name = user.fullName || user.name || user.id;
                if (tries > 1) {
                    name = name + " " + tries;
                }
                team = await this.teamDB.createTeam(user.id, name);
            } catch (err) {
                if (err instanceof ResponseError) {
                    if (err.code === ErrorCodes.CONFLICT) {
                        continue;
                    }
                }
                throw err;
            }
        }
        if (!team) {
            throw new ResponseError(ErrorCodes.CONFLICT, "Could not create team for user.", { userId: user.id });
        }

        const projects = await this.projectDB.findUserProjects(user.id);
        log.info(ctx, "Migrating projects.", { teamId: team.id, projects: projects.map((p) => p.id) });
        for (const project of projects) {
            project.teamId = team.id;
            project.userId = "";
            await this.projectDB.storeProject(project);
        }

        const conn = await this.typeorm.getConnection();
        const oldAttribution = "user:" + user.id;
        const newAttribution = "team:" + team.id;

        let result = await conn.query(
            "UPDATE d_b_workspace_instance SET usageAttributionId = ? WHERE usageAttributionId = ?",
            [newAttribution, oldAttribution],
        );
        log.info(ctx, "Migrated workspace instances.", { teamId: team.id, result });

        result = await conn.query("UPDATE d_b_usage SET attributionId = ? WHERE attributionId = ?", [
            newAttribution,
            oldAttribution,
        ]);
        log.info(ctx, "Migrated usage data.", { teamId: team.id, result });

        result = await conn.query("UPDATE d_b_cost_center SET id = ? WHERE id = ?", [newAttribution, oldAttribution]);
        log.info(ctx, "Migrated cost center data.", { teamId: team.id, result });

        result = await conn.query("UPDATE d_b_stripe_customer SET attributionid = ? WHERE attributionid = ?", [
            newAttribution,
            oldAttribution,
        ]);
        log.info(ctx, "Migrated stripe customer data.", { teamId: team.id, result });
        return team;
    }

    async needsMigration(user: User): Promise<boolean> {
        const teams = await this.teamDB.findTeamsByUser(user.id);
        if (teams.length === 0) {
            return true;
        }
        const projects = await this.projectDB.findUserProjects(user.id);
        if (projects.length > 0) {
            return true;
        }
        const instances = await (
            await this.typeorm.getConnection()
        ).query("SELECT * from d_b_workspace_instance where usageAttributionId=?", ["user:" + user.id]);
        if (instances.length > 0) {
            return true;
        }
        return false;
    }
}
