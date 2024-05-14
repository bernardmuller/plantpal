import {httpRequest} from "$lib/http/httpRequest";
import type {Plant} from "../../types";

export async function getAllPlants() {
  return httpRequest<Plant[], undefined>("http://localhost:8080/plants");
}