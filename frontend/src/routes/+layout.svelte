<script>
    import { initAuthContext, setUser } from "$lib";
    import Navigation from "$lib/Navigation.svelte";

    /** @type {import('./$types').LayoutData} */
    export let data;

    initAuthContext();

    $: if (data.user) {
        setUser(data.user);
    } else {
        setUser(null);
    }
</script>

{#if data.user}<Navigation />{/if}
<main>
    <slot />
</main>
