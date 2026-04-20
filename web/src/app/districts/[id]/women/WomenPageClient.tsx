"use client";

import { useRouter, useSearchParams } from "next/navigation";
import { usePageParam } from "@/components/pagination/usePageParam";
import WomenSection from "./WomenSection";

type Props = {
  districtId: string;
};

export default function WomenPageClient({ districtId }: Props) {
  const router = useRouter();
  const searchParams = useSearchParams();
  const { page, goToPage } = usePageParam();

  const bloodTypes = searchParams.getAll("blood_type");
  const ageRanges = searchParams.getAll("age_range");

  const onSearch = (newBloodTypes: string[], newAgeRanges: string[]) => {
    const params = new URLSearchParams();
    params.set("page", "1");
    newBloodTypes.forEach((bt) => params.append("blood_type", bt));
    newAgeRanges.forEach((ar) => params.append("age_range", ar));
    router.push(`?${params.toString()}`);
  };

  return (
    <WomenSection
      districtId={districtId}
      page={page}
      onPageChange={goToPage}
      bloodTypes={bloodTypes}
      ageRanges={ageRanges}
      onSearch={onSearch}
    />
  );
}
