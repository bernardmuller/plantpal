import {writable} from 'svelte/store'

let stored;
if (typeof window !== 'undefined') {
  stored = localStorage.getItem("pp_token")
}

export const token = writable(stored || 'Hello, World!')

token.subscribe((value) => {
  if (typeof window !== 'undefined') {
    localStorage.setItem("pp_token", value)
  }
})