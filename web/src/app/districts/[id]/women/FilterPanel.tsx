"use client";

import { useCallback } from "react";
import CommonFilterPanel from "@/components/conditions/WomanFilterPanel";
import { fetchDistrictWomanCount } from "@/api/woman";

type Props = {
  districtId: string;
  onSearch?: (bloodTypes: string[], ageRanges: string[]) => void;
  initialBloodTypes?: string[];
  initialAgeRanges?: string[];
};

export default function FilterPanel({ districtId, onSearch, initialBloodTypes, initialAgeRanges }: Props) {
  const fetchSearchCount = useCallback(
    (bloodTypes: string[], ageRanges: string[]) =>
      fetchDistrictWomanCount(districtId, bloodTypes, ageRanges).then((d) => d.total),
    [districtId]
  );

  return (
    <CommonFilterPanel
      fetchSearchCount={fetchSearchCount}
      onSearch={onSearch}
      initialBloodTypes={initialBloodTypes}
      initialAgeRanges={initialAgeRanges}
    />
  );
}
