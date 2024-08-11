import { extractRoundIdFromUrl } from "$lib";
import { error } from "@sveltejs/kit";

/** @type {import("./$types").LayoutServerLoad}*/
export const load = async ({ url }) => {
    const roundId = extractRoundIdFromUrl(url)
    console.log("round id in layout: ", roundId)
    if (roundId === undefined) {
        error(400, "No round provided")
    }
    const response = await fetch(`http://localhost:8800/rounds/${roundId}`)
    if (response.ok) {
        /**
         * @type {import("$lib/types").Round}
         */
        const round = await response.json()
        return { round }
    }
    return { round: null }
}