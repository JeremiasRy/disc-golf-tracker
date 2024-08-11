// src/lib/roundContext.js
import { writable } from 'svelte/store';
import { setContext, getContext } from 'svelte';
export const ROUND_CONTEXT = "round"

/**
 * @type {import("svelte/store").Writable<import("$lib/types").Round | null>}
 */
const roundStore = writable(null);

/**
 * @param {import("$lib/types").Round | null} round 
 */
export function setRoundContext(round) {
    roundStore.set(round);

}

export function initRoundContext() {
    setContext(ROUND_CONTEXT, roundStore);
}

/**
 * @returns {import("svelte/store").Writable<import("$lib/types").Round | null>}
 */
export function useRoundContext() {
    return getContext(ROUND_CONTEXT);
}