<script>
    import { goto } from "$app/navigation";
    import { useRoundContext } from "$lib";

    const roundStore = useRoundContext();
    /**
     * @type {import("$lib/types").Round | null}
     */
    let round;
    roundStore.subscribe((value) => {
        round = value;
    });

    const handleBegin = () => {
        goto(`/app/rounds/${round?.ID}/play/1`);
    };
</script>

{#if !round}
    <h1>Loading...</h1>
{:else}
    <h1>{round.course.name}</h1>
    <h4>Players Invited</h4>
    <ul>
        {#each round.cards as card}
            <li>{card.user.name}</li>
        {/each}
    </ul>
    <h4>Invite more players</h4>
    <form>
        <input />
        <input type="submit" value="Invite" />
    </form>
    <br />
    <button on:click={handleBegin}>Begin!</button>
{/if}
