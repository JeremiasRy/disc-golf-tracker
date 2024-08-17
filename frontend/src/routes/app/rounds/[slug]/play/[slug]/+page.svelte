<script>
    import { page } from "$app/stores";
    import { error } from "@sveltejs/kit";

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

    let currentScores = scores
        .filter((score) => score.hole_id === holeId)
        .map((score) => ({
            score,
            user: round.cards.find((card) => card.ID === score.scorecard_id)
                ?.user,
        }));

    /**
     * @param newScore {import("$lib/types.js").Score}
     */
    function updateScore(newScore) {
        currentScores = currentScores.map(({ user, score }) => ({
            user,
            score: score.ID !== newScore.ID ? score : newScore,
        }));
    }

    /**
     * @param {MouseEvent & { currentTarget: EventTarget & HTMLButtonElement; }} event
     */
    async function updateStrokes({ currentTarget }) {
        const [newStrokes, scoreId] = currentTarget.value.split("|");
        const updateStrokesRequest = {
            method: "PATCH",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ new_strokes: Number(newStrokes) }),
        };
        const response = await fetch(
            `http://localhost:8800/scores/${scoreId}`,
            updateStrokesRequest,
        );

        if (response.ok) {
            const newScore = await response.json();
            updateScore(newScore);
            return;
        }
        console.log("We failed :/");
    }
    /**
     * @param {MouseEvent & { currentTarget: EventTarget & HTMLButtonElement; }} event
     */
    async function updatePenalties({ currentTarget }) {
        const [newPenalties, scoreId] = currentTarget.value.split("|");
        const updatePenaltiesRequest = {
            method: "PATCH",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ new_penalties: Number(newPenalties) }),
        };
        const response = await fetch(
            `http://localhost:8800/scores/${scoreId}`,
            updatePenaltiesRequest,
        );

        if (response.ok) {
            const newScore = await response.json();
            updateScore(newScore);
            return;
        }
    }

    function getNextHoleHref() {
        const currentPath = $page.url.pathname;
        const segments = currentPath.split("/");

        const currentId = Number(segments.pop());
        const nextId = currentId + 1;

        return `${segments.join("/")}/${nextId}`;
    }

    function getPreviousHoleHref() {
        const currentPath = $page.url.pathname;
        const segments = currentPath.split("/");

        const currentId = Number(segments.pop());
        const nextId = currentId - 1;

        return `${segments.join("/")}/${nextId}`;
    }
</script>

<h2>Hole {nthHole} Par {par}</h2>
{#each currentScores as { score, user }}
    <h4>{user?.name}</h4>
    <div class="score-input-wrapper">
        <div class="score-input-wrapper__score-input">
            <button
                on:click={updateStrokes}
                value={`${score.strokes + 1}|${score.ID}`}>+</button
            >
            {score.strokes} Strokes
            <button
                on:click={updateStrokes}
                value={`${score.strokes - 1}|${score.ID}`}
                disabled={score.strokes === 0}>-</button
            >
        </div>
        <div class="score-input-wrapper__divider" />
        <div class="score-input-wrapper__score-input">
            <button
                on:click={updatePenalties}
                value={`${score.penalties + 1}|${score.ID}`}>+</button
            >
            {score.penalties} Penalties
            <button
                on:click={updatePenalties}
                value={`${score.penalties - 1}|${score.ID}`}
                disabled={score.penalties === 0}>-</button
            >
        </div>
    </div>
    <form method="POST" action="?/redirect">
        <input type="hidden" value={getNextHoleHref()} name="url" />
        <button type="submit">Next Hole</button>
    </form>
    <form method="POST" action="?/redirect">
        <input type="hidden" value={getPreviousHoleHref()} name="url" />
        <button type="submit" disabled={nthHole === "1"}>Previous Hole</button>
    </form>
{/each}
