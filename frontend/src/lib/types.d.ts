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
    nth_hole: number
    par: number
    ID: number
}

export type Round = {
    ID: number
    course: Course
    cards: ScoreCard[]
}

export type ScoreCard = {
    ID: number
    user: User
    scores: Score[]
}

export type Score = {
    ID: number
    hole_id: number
    scorecard_id: number
    strokes: number
    penalties: number
}