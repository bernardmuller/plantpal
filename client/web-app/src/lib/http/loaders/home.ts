import { getAllPlants } from "../client/getAllPlants";

export default async function homeLoader() {
	const plantsResponse = await getAllPlants();
	return { plants: plantsResponse.data ?? null };
}
