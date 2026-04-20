"use client";

export type AgeRange = {
  min: number;
  max: number;
};

const AGE_RANGES: AgeRange[] = [
  { min: 20, max: 25 },
  { min: 26, max: 30 },
  { min: 31, max: 35 },
  { min: 36, max: 40 },
];

const toKey = (range: AgeRange) => `${range.min}-${range.max}`;

type Props = {
  selected: AgeRange[];
  onChange: (conditionAgeRanges: AgeRange[]) => void;
};

export default function AgeFilter({ selected, onChange }: Props) {
  const isSelected = (range: AgeRange) =>
    selected.some((r) => r.min === range.min && r.max === range.max);

  const handleChange = (range: AgeRange, checked: boolean) => {
    const conditionAgeRanges = checked
      ? [...selected, range]
      : selected.filter((r) => r.min !== range.min || r.max !== range.max);

    onChange(conditionAgeRanges);
  };

  return (
    <div>
      <p>年齢で絞り込み：</p>
      {AGE_RANGES.map((range) => (
        <label key={toKey(range)}>
          <input
            type="checkbox"
            value={toKey(range)}
            checked={isSelected(range)}
            onChange={(e) => handleChange(range, e.target.checked)}
          />
          {range.min}〜{range.max}歳
        </label>
      ))}
    </div>
  );
}
