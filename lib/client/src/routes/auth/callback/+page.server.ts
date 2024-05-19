import {redirect} from '@sveltejs/kit';

export const load = ((context) => {

  const token = context.request.url.toString().split("?")[1].split("=")[1]
  if (!token) {
    redirect(307, "/auth/login?error=invalid_token")
  }
  return {
    token: token
  }
})