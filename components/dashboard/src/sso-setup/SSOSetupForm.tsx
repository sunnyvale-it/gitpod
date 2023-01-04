/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { FunctionComponent, useCallback, useReducer, useState } from "react";
import { TextInput } from "../components/forms/TextInput";
import { oidcService } from "../service/public-api";
import Alert from "../components/Alert";

// TODO: Use a type already in gitpod-protocol
type SSOConfig = {
    clientID: string;
    clientSecret: string;
    authEndpoint: string;
};

type SSOSetupFormProps = {
    token: string;
};

export const SSOSetupForm: FunctionComponent<SSOSetupFormProps> = ({ token }) => {
    const [showSaveError, setShowSaveError] = useState(false);
    const [saving, setSaving] = useState(false);
    const [config, dispatch] = useReducer(
        (state: SSOConfig, action: Partial<SSOConfig>) => ({
            ...state,
            ...action,
        }),
        {
            clientID: "",
            clientSecret: "",
            authEndpoint: "",
        },
    );

    const updateClientID = useCallback((val) => dispatch({ clientID: val }), []);
    const updateClientSecret = useCallback((val) => dispatch({ clientSecret: val }), []);
    const updateAuthEndpoint = useCallback((val) => dispatch({ authEndpoint: val }), []);

    const handleSave = useCallback(
        async (e) => {
            e.preventDefault();
            setSaving(true);

            console.log("saving config", config);

            try {
                const { config: newConfig } = await oidcService.createClientConfig({
                    config: {
                        oauth2Config: {
                            clientId: config.clientID,
                            clientSecret: config.clientSecret,
                            authorizationEndpoint: config.authEndpoint,
                        },
                    },
                });

                console.log("saved config", newConfig);
            } catch (e) {
                console.log("Error saving SSO Config", e);
                setShowSaveError(true);
            }

            setSaving(false);
        },
        [config, token],
    );

    return (
        <div>
            {showSaveError && <Alert type="error">Sorry, there was an error saving your SSO Configuration</Alert>}
            <form onSubmit={handleSave}>
                <TextInput label="Client ID" value={config.clientID} disabled={saving} onChange={updateClientID} />
                <TextInput
                    label="Client Secret"
                    value={config.clientSecret}
                    disabled={saving}
                    onChange={updateClientSecret}
                />
                <TextInput
                    label="Authorization Endpoint"
                    value={config.authEndpoint}
                    disabled={saving}
                    onChange={updateAuthEndpoint}
                />

                <div className="mt-4">
                    <button disabled={saving}>Save</button>
                </div>
            </form>
        </div>
    );
};
