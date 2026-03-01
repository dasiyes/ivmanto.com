import axios from 'axios'

// Define the structure of the ideas we expect from the backend
export interface Idea {
  title: string
  summary: string
}

/**
 * Calls the backend to generate ideas based on a topic.
 * @param topic The user-provided topic.
 * @returns A promise that resolves to an array of ideas.
 */
export async function generateInspirationIdeas(topic: string): Promise<Idea[]> {
  try {
    const response = await axios.post<Idea[]>('/api/generate-ideas', { topic })
    return response.data
  } catch (error) {
    throw new Error('Failed to generate ideas. Please try again later.')
  }
}
