import { BASE_URL, fetchJSON } from "./base";
import type { ImmediateAvailableWomanListResponse } from "@/interfaces/immediateAvailableWoman/list";

export async function fetchImmediateAvailableWomen(params: {
  page?: number;
  prefectureId?: number;
  districtId?: number;
  businessTypes?: string[];
  bloodTypes?: string[];
  ageRanges?: string[];
}, options?: RequestInit): Promise<ImmediateAvailableWomanListResponse> {
  const query = new URLSearchParams();
  if (params.page) query.set("page", String(params.page));
  if (params.prefectureId) query.set("prefecture_id", String(params.prefectureId));
  if (params.districtId) query.set("district_id", String(params.districtId));
  if (params.businessTypes && params.businessTypes.length > 0) query.set("business_type", params.businessTypes.join(","));
  if (params.bloodTypes && params.bloodTypes.length > 0) query.set("blood_type", params.bloodTypes.join(","));
  if (params.ageRanges && params.ageRanges.length > 0) query.set("age_range", params.ageRanges.join(","));
  return fetchJSON<ImmediateAvailableWomanListResponse>(`${BASE_URL}/immediate_available_women?${query.toString()}`, options);
}
