import apiRoutes from "../routes";
import { httpRequest } from "../httpClient";
import { z } from "zod";

const plantSchema = z.object({
	ID: z.string(),
	Common: z.string(),
	Family: z.string(),
	CreatedAt: z.string(),
	UpdatedAt: z.string(),
	Latin: z.string().nullable(),
	Category: z.string().nullable(),
	Origin: z.string().nullable(),
	Climate: z.string().nullable(),
	Tempmax: z.string().nullable(),
	Tempmin: z.string().nullable(),
	Ideallight: z.string().nullable(),
	Toleratedlight: z.string().nullable(),
	Watering: z.string().nullable(),
	Insects: z.string().nullable(),
	Diseases: z.string().nullable(),
	Soil: z.string().nullable(),
	Repotperiod: z.string().nullable(),
	Use: z.string().nullable(),
});

export type Plant = z.infer<typeof plantSchema>;

export async function getAllPlants() {
	return await httpRequest<Plant[], void>(apiRoutes.getAllPlants, "GET");
}
