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

    const { ID: holeId, par } = currentHole;

    const currentScores = scores
        .filter((score) => score.hole_id === holeId)
        .map((score) => ({
            score,
            user: round.cards.find((card) => card.ID === score.scorecard_id)
                ?.user,
        }));
</script>

<h2>Hole {nthHole} Par {par}</h2>
{#each currentScores as { score, user }}
    <h4>{user?.name}</h4>
    <p>
        {score.strokes} Strokes | {score.penalties} Penalties
    </p>
{/each}
