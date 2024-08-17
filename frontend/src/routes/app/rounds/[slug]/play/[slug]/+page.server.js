import { error, redirect } from "@sveltejs/kit"

/**
 * @type {import("./$types").PageServerLoad}
 */
export const load = async ({ url, parent }) => {
    const { round } = await parent()
    const nthHole = url.pathname.split("/").pop()

    if (nthHole === undefined || !round) {
        error(400, "A very bad request")
    }

    const hole = round.course.holes.find(hole => hole.nth_hole === Number(nthHole))

    if (!hole) {
        error(400, "a not so bad request")
    }

    const scores = await Promise.all(round.cards.map((card) => createScore(hole.ID, card.ID)))

    return { nthHole, scores }
}

/**
 * Create or fetch score
 * @param {number} holeId 
 * @param {number} scoreCardId
 * @returns {Promise<import("$lib/types").Score>} 
 */
async function createScore(holeId, scoreCardId) {
    const scoreRequest = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ hole_id: holeId, score_card_id: scoreCardId })
    }
    const response = await fetch("http://localhost:8800/scores", scoreRequest)
    if (response.ok || response.status === 304) {
        /**
         * @type {import("$lib/types").Score}
         */
        const score = await response.json()
        return score
    }

    error(400, "Failed to create or fetch score")
}
/** @type {import("./$types").Actions} */
export const actions = {
    redirect: async ({ request }) => {
        const data = await request.formData()
        const url = String(data.get("url"))
        redirect(302, url)
    }
}