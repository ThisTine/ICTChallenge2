export interface Question {
	mode: Mode
	topics: Topic[]
}

export interface Topic {
	id?: number
	title: string
	cards: Card[]
}

export interface Card {
	id: number
	score: number
	opened: boolean
}

export interface OpenQuestion {
	topic_id: number
	question_id: number
	title: string
	imageUrl?: string
	bonus: boolean
}

export interface QuestionDetail {
	title: string
	imageUrl?: string
}

export enum Mode {
	RULE = 'rule',
	TOPIC = 'topic',
}
