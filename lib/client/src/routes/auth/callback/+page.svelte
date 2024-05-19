<!--export const prerender = true;-->
<script>
    import {token} from "../../../lib/store"

    export let data

    token.set(data.token)

    if (typeof window !== 'undefined') {
        window.location.href = '/dashboard'
    }
</script>


<div>redirecting...</div>