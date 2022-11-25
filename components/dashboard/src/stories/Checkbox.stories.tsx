/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { ComponentMeta, ComponentStory } from "@storybook/react";
import Checkbox from "../components/CheckBox";

export default {
    title: "Components/Checkbox",
    component: Checkbox,
    // More on argTypes: https://storybook.js.org/docs/react/api/argtypes
    argTypes: {
        backgroundColor: { control: "color" },
    },
} as ComponentMeta<typeof Checkbox>;

const Template: ComponentStory<typeof Checkbox> = (args) => <Checkbox {...args} />;

export const Default = Template.bind({});
Default.args = {
    title: "",
    desc: "",
    disabled: false,
};

export const Disabled = Template.bind({});
Disabled.args = {
    title: "",
    desc: "",
    disabled: true,
};
