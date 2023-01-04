/**
 * Copyright (c) 2023 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { FunctionComponent, memo, useCallback } from "react";
import { useId } from "../../hooks/useId";

type TextInputProps = {
    label: string;
    value: string;
    id?: string;
    placeholder?: string;
    disabled?: boolean;
    required?: boolean;
    onChange: (newValue: string) => void;
};

export const TextInput: FunctionComponent<TextInputProps> = memo(
    ({ label, value, id, placeholder, disabled = false, required = false, onChange }) => {
        const maybeId = useId();
        const elementId = id || maybeId;
        console.log(label, elementId);

        const handleChange = useCallback(
            (e) => {
                onChange(e.target.value);
            },
            [onChange],
        );

        return (
            <div className="mt-4">
                <label className="text-sm font-semibold text-gray-600 dark:text-gray-400" htmlFor={elementId}>
                    {label}
                </label>
                <input
                    id={elementId}
                    className="w-full max-w-lg"
                    value={value}
                    onChange={handleChange}
                    type="text"
                    placeholder={placeholder}
                    disabled={disabled}
                    required={required}
                />
            </div>
        );
    },
);
