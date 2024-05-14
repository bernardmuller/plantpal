import {getAllPlants} from "$lib/api/getAllPlants";
import {getCurrentUser} from "$lib/api/getCurrentUser";

const plantsResponse = await getAllPlants()
const currentUserResponse = await getCurrentUser()


export const load = (() => {
  console.log("loading plants page")
  console.log(currentUserResponse)
  return {
    currentUser: structuredClone(currentUserResponse.ok ? currentUserResponse.data : null),
    plants: structuredClone(plantsResponse.ok ? plantsResponse.data : [])
  };
})