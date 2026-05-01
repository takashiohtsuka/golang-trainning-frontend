import { BASE_URL, fetchJSON } from "./base";
import type { PrefectureListResponse } from "@/interfaces/prefecture/list";

export async function fetchPrefectures(options?: RequestInit): Promise<PrefectureListResponse> {
  return fetchJSON<PrefectureListResponse>(`${BASE_URL}/prefectures`, options);
}
