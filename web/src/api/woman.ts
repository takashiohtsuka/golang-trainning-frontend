import type { DistrictWomenResponse } from "@/interfaces/district/womanList";
import type { DistrictWomanCountResponse } from "@/interfaces/district/womanCount";
import { BASE_URL } from "./base";

export async function fetchDistrictWomen(
  districtId: string,
  bloodTypes: string[] = [],
  ageRanges: string[] = [],
  page: number = 1
): Promise<DistrictWomenResponse> {
  const params = new URLSearchParams();
  params.set("page", String(page));
  bloodTypes.forEach((bt) => params.append("blood_type", bt));
  ageRanges.forEach((r) => params.append("age_range", r));
  const res = await fetch(`${BASE_URL}/districts/${districtId}/women?${params.toString()}`);
  if (!res.ok) throw new Error(`${res.status}`);
  return res.json() as Promise<DistrictWomenResponse>;
}

export async function fetchDistrictWomanCount(
  districtId: string,
  bloodTypes: string[],
  ageRanges: string[]
): Promise<DistrictWomanCountResponse> {
  const params = new URLSearchParams();
  bloodTypes.forEach((bt) => params.append("blood_type", bt));
  ageRanges.forEach((r) => params.append("age_range", r));
  const res = await fetch(`${BASE_URL}/districts/${districtId}/search-woman-count?${params.toString()}`);
  if (!res.ok) throw new Error(`${res.status}`);
  return res.json() as Promise<DistrictWomanCountResponse>;
}
