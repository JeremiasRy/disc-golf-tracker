export type User = {
    name: string
    email: string
    ID: string
}

export type Course = {
    ID: string
    name: string
    holes: Hole[]
}

export type Hole = {
    nthHole: number
    par: number
    ID: number
}

export type Round = {
    ID: number
    ScoreCards: ScoreCard[]
}

export type ScoreCard = {
    ID: number
    scores: Score[]
}

export type Score = {
    ID: number
    strokes: number
    penalties: number
}