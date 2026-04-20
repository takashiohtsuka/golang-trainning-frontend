"use client";

import FilterPanel from "./FilterPanel";
import WomenList from "./WomenList";

type Props = {
  districtId: string;
  page: number;
  onPageChange: (page: number) => void;
  bloodTypes: string[];
  ageRanges: string[];
  onSearch: (bloodTypes: string[], ageRanges: string[]) => void;
};

export default function WomenSection({ districtId, page, onPageChange, bloodTypes, ageRanges, onSearch }: Props) {
  return (
    <>
      <FilterPanel
        districtId={districtId}
        onSearch={onSearch}
        initialBloodTypes={bloodTypes}
        initialAgeRanges={ageRanges}
      />
      <WomenList
        districtId={districtId}
        bloodTypes={bloodTypes}
        ageRanges={ageRanges}
        page={page}
        onPageChange={onPageChange}
      />
    </>
  );
}
