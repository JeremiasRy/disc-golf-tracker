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