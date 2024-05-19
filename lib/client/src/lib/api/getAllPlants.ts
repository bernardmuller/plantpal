import {httpRequest} from "$lib/http/httpRequest";
import type {Plant} from "../../types";

export async function getAllPlants() {
  return httpRequest<Plant[], undefined>("http://localhost:8080/plants?user=123");
//   return await fetch("http://localhost:8080/plants?user=123", {
// // method: "GET",
// //     headers: {
// //       "Access-Control-Allow-Origin": "*",
// //       "Access-Control-Allow-Credentials": "true",
// //       "Content-Type": "application/json",
// //     },
// //     credentials: 'include'
//   })
}