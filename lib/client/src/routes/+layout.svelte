<script>
    import "../styles/globals.css";
    import {QueryClient, QueryClientProvider} from '@tanstack/svelte-query'

    const queryClient = new QueryClient()
</script>

<QueryClientProvider client={queryClient}>
    <main class="w-full flex flex-col items-center bg-background dark">
        <section class="w-full max-w-[1444px] p-8">
            <slot/>
        </section>
    </main>
</QueryClientProvider>
