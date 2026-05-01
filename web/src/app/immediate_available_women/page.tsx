export const dynamic = "force-dynamic";

import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import ImmediateAvailableWomenPageClient from "./ImmediateAvailableWomenPageClient";
import { fetchImmediateAvailableWomen } from "@/api/immediateAvailableWoman";
import { fetchPrefectures } from "@/api/prefecture";
import { fetchBusinessTypes } from "@/api/businessType";

export default async function Page() {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { staleTime: 30 * 1000 } },
  });

  const [prefectureData, businessTypeData] = await Promise.all([
    fetchPrefectures({ next: { revalidate: 60 * 60 } }),
    fetchBusinessTypes({ next: { revalidate: 60 * 60 } }),
  ]);

  await queryClient.prefetchQuery({
    queryKey: ["immediateAvailableWomen", 1, undefined, undefined, [], [], []],
    queryFn: () => fetchImmediateAvailableWomen({ page: 1 }, { next: { revalidate: 30 } }),
  });

  return (
    <main>
      <h1>すぐ予約できる女性一覧</h1>
      <HydrationBoundary state={dehydrate(queryClient)}>
        <ImmediateAvailableWomenPageClient
          prefectures={prefectureData.prefectures}
          businessTypes={businessTypeData.business_types}
        />
      </HydrationBoundary>
    </main>
  );
}
