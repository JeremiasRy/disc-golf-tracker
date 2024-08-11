<script>
    import { error, json } from "@sveltejs/kit";
    export let data;
    const { round, nthHole, scores } = data;

    if (!round) {
        error(400, "No round defined, something went wrong");
    }

    /**
     * @type {import("$lib/types.js").Hole | undefined}
     */
    let currentHole = round.course.holes.find(
        (hole) => hole.nth_hole === Number(nthHole),
    );

    if (!currentHole) {
        error(404, "Ended up in a hole that doesn't exist");
    }
</script>

<h2>Hole {nthHole} Par {currentHole.par}</h2>
{#each scores as score}
    {JSON.stringify(score)}
{/each}
