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
    console.error('API Error in generateInspirationIdeas:', error)
    // Re-throw a new, more generic error to be handled by the calling component.
    // This prevents the function from implicitly returning `undefined` and causing a type error.
    throw new Error('Failed to generate ideas. Please try again later.')
  }
}
