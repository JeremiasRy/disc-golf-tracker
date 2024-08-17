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

    const { par } = currentHole;

    let currentScores = scores.map((score) => ({
        score,
        user: round.cards.filter((card) => card.ID === score.scorecard_id)[0]
            .user,
    }));

    let scoreDisplay = round.cards.map((card) => ({
        user: card.user,
        scores: card.scores.map((score) => ({
            par: round.course.holes.filter(
                (hole) => hole.ID === score.hole_id,
            )[0].par,
            score,
        })),
    }));

    /**
     * @param newScore {import("$lib/types.js").Score}
     */
    function updateScore(newScore) {
        currentScores = currentScores.map(({ user, score }) => ({
            user,
            score: score.ID !== newScore.ID ? score : newScore,
        }));

        scoreDisplay = scoreDisplay.map(({ user, scores }) => ({
            user,
            scores: scores.map(({ par, score }) => ({
                par,
                score: score.ID === newScore.ID ? newScore : score,
            })),
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

{#each scoreDisplay as { user, scores }}
    <h2>{user.name}</h2>
    <div>
        {#each scores as { par, score }}
            <div>{score.strokes + score.penalties}</div>
        {/each}
    </div>
{/each}
<h2>Hole {nthHole} Par {par}</h2>

{#each currentScores as { score, user }}
    <h4>{user?.name} | {score.penalties + score.strokes}</h4>
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
        <button
            type="submit"
            disabled={Number(nthHole) === round.course.holes.length}
            >Next Hole</button
        >
    </form>
    <form method="POST" action="?/redirect">
        <input type="hidden" value={getPreviousHoleHref()} name="url" />
        <button type="submit" disabled={nthHole === "1"}>Previous Hole</button>
    </form>
{/each}
