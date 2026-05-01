import { BASE_URL, fetchJSON } from "./base";
import type { DistrictWomenResponse } from "@/interfaces/district/womanList";
import type { DistrictWomanCountResponse } from "@/interfaces/district/womanCount";

export async function fetchDistrictWomen(
  districtId: string,
  bloodTypes: string[] = [],
  ageRanges: string[] = [],
  page: number = 1,
  options?: RequestInit
): Promise<DistrictWomenResponse> {
  const params = new URLSearchParams();
  params.set("page", String(page));
  if (bloodTypes.length > 0) params.set("blood_type", bloodTypes.join(","));
  if (ageRanges.length > 0) params.set("age_range", ageRanges.join(","));
  return fetchJSON<DistrictWomenResponse>(`${BASE_URL}/districts/${districtId}/women?${params.toString()}`, options);
}

export async function fetchDistrictWomanCount(
  districtId: string,
  bloodTypes: string[],
  ageRanges: string[]
): Promise<DistrictWomanCountResponse> {
  const params = new URLSearchParams();
  if (bloodTypes.length > 0) params.set("blood_type", bloodTypes.join(","));
  if (ageRanges.length > 0) params.set("age_range", ageRanges.join(","));
  return fetchJSON<DistrictWomanCountResponse>(`${BASE_URL}/districts/${districtId}/search-woman-count?${params.toString()}`);
}
