import apiRoutes from "../routes";
import { httpRequest } from "../httpClient";
import { Plant } from "./getAllPlants";

export async function getPlantById(plantId: string) {
	return await httpRequest<Plant, void>(
		apiRoutes.getPlantById(plantId),
		"GET"
	);
}
