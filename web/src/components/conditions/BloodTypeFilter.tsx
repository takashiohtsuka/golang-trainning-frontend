"use client";

const BLOOD_TYPES = ["A", "B", "O", "AB"] as const;
export type BloodType = (typeof BLOOD_TYPES)[number];

type Props = {
  selected: BloodType[];
  onChange: (conditionBloodTypes: BloodType[]) => void;
};

export default function BloodTypeFilter({ selected, onChange }: Props) {
  const handleChange = (bloodType: BloodType, checked: boolean) => {
    const conditionBloodTypes = checked
      ? [...selected, bloodType]
      : selected.filter((bt) => bt !== bloodType);

    onChange(conditionBloodTypes);
  };

  return (
    <div>
      <p>血液型で絞り込み：</p>
      {BLOOD_TYPES.map((bt) => (
        <label key={bt}>
          <input
            type="checkbox"
            value={bt}
            checked={selected.includes(bt)}
            onChange={(e) => handleChange(bt, e.target.checked)}
          />
          {bt}型
        </label>
      ))}
    </div>
  );
}
