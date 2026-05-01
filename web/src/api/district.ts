import { BASE_URL, fetchJSON } from "./base";
import type { DistrictListResponse } from "@/interfaces/district/list";

export async function fetchDistricts(prefectureId: number, options?: RequestInit): Promise<DistrictListResponse> {
  return fetchJSON<DistrictListResponse>(`${BASE_URL}/prefectures/${prefectureId}/districts`, options);
}
