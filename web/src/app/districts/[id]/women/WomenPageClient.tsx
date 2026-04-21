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

  const bloodTypes = searchParams.get("blood_type")?.split(",").filter(Boolean) ?? [];
  const ageRanges = searchParams.get("age_range")?.split(",").filter(Boolean) ?? [];

  const onSearch = (newBloodTypes: string[], newAgeRanges: string[]) => {
    const params = new URLSearchParams();
    params.set("page", "1");
    if (newBloodTypes.length > 0) params.set("blood_type", newBloodTypes.join(","));
    if (newAgeRanges.length > 0) params.set("age_range", newAgeRanges.join(","));
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
