import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import WomenPageClient from "./WomenPageClient";
import { fetchDistrictWomen } from "@/api/woman";
import type { DistrictWomenResponse } from "@/interfaces/district/womanList";

type Props = {
  params: Promise<{ id: string }>;
};

export default async function Page({ params }: Props) {
  const { id } = await params;

  const queryClient = new QueryClient({
    defaultOptions: { queries: { staleTime: 30 * 1000 } },
  });

  await queryClient.prefetchQuery({
    queryKey: ["districtWomen", id, [], [], 1],
    queryFn: () => fetchDistrictWomen(id, [], [], 1, { next: { revalidate: 30 } }),
  });

  const womenData = queryClient.getQueryData<DistrictWomenResponse>(["districtWomen", id, [], [], 1]);
  queryClient.setQueryData(["districtWomanCount", id, [], []], womenData?.total ?? 0);

  return (
    <main>
      <h1>District {id} の女性一覧</h1>
      <HydrationBoundary state={dehydrate(queryClient)}>
        <WomenPageClient districtId={id} />
      </HydrationBoundary>
    </main>
  );
}
