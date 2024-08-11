import { error, redirect } from '@sveltejs/kit';
import { nodeModuleNameResolver } from 'typescript';

/** @type {import('./$types').PageServerLoad} */
export const load = async ({ params, fetch }) => {
    const courseId = params.slug;
    const response = await fetch(`http://localhost:8800/courses/${courseId}`);

    if (response.ok) {
        /**
         * @type {import("$lib/types").Course}
         */
        const course = await response.json()
        return { course }
    }
    error(404, "Not found")
}

export const actions = {
    /**
    * Start a new round and redirect to the new round's page.
    * 
    * @param {import('@sveltejs/kit').RequestEvent} event 
    * @returns {Promise<void>} Redirect to the new round
    */
    startNewRound: async ({ request }) => {
        const data = await request.formData()
        const courseId = Number(data.get('courseId'));
        const userId = Number(data.get('userId'))

        const roundRequest = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ course_id: courseId })
        }
        const response = await fetch("http://localhost:8800/rounds", roundRequest)

        if (response.ok) {
            const round = await response.json()
            const scorecardRequest = {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ user_id: userId, round_id: round.ID })
            }
            await fetch("http://localhost:8800/scorecards", scorecardRequest) // we dont care if this fails, we can just ask to create again

            redirect(303, `/app/rounds/${round.ID}`)
        } else {
            throw error(response.status, 'Failed to create new round');
        }
    }
};