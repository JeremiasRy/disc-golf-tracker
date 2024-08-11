<script>
    import { initRoundContext, setRoundContext } from "$lib";

    /** @type {import('./$types').LayoutData} */
    export let data;

    initRoundContext();

    $: if (data.round) {
        setRoundContext(data.round);
    } else {
        setRoundContext(null);
    }
</script>

<slot />
