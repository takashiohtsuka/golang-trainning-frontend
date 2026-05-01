"use client";

import type { BusinessTypeItem } from "@/interfaces/businessType/list";

type Props = {
  businessTypes: BusinessTypeItem[];
  selected: string[];
  onChange: (selected: string[]) => void;
};

export default function BusinessTypeFilter({ businessTypes, selected, onChange }: Props) {
  const handleChange = (code: string, checked: boolean) => {
    onChange(checked ? [...selected, code] : selected.filter((v) => v !== code));
  };

  return (
    <div>
      <p>業種：</p>
      {businessTypes.map((bt) => (
        <label key={bt.code}>
          <input
            type="checkbox"
            value={bt.code}
            checked={selected.includes(bt.code)}
            onChange={(e) => handleChange(bt.code, e.target.checked)}
          />
          {bt.code}
        </label>
      ))}
    </div>
  );
}
