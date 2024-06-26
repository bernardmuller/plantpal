import { getPlantById } from "../client/getPlantById";

export default async function plantDetailLoader({
	request,
}: {
	request: Request;
}) {
	const url = new URL(request.url);
	const plantId = url.pathname.split("/")[2];
	const plantsResponse = await getPlantById(plantId);
	return { plant: plantsResponse.data ?? null };
}
