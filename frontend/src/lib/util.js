/**
 * @param {URL} url 
 * @returns round id as a string or empty string if roundId is not present
 */
export const extractRoundIdFromUrl = ({ pathname }) => {
    return pathname.split("/")[3];
}