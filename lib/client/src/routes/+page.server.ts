import {getAllPlants} from "$lib/api/getAllPlants";

const plantsResponse = await getAllPlants()

export const load = (() => {
  return {
    plants: structuredClone(plantsResponse.ok ? plantsResponse.data : [])
  };
})