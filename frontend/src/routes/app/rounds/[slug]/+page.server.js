import { error } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export const load = async ({ params, fetch }) => {
    const roundId = params.slug;
    const response = await fetch(`http://localhost:8800/rounds/${roundId}`);

    if (response.ok) {
        /**
         * @type {import("../../../../lib/types").Course}
         */
        const round = await response.json()
        return { round }
    }
    error(404, "Not found")
}