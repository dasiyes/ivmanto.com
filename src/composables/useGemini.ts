import { ref } from 'vue'

// --- Define specific types for better type safety ---
interface GeminiPart {
  text: string
}

interface GeminiContent {
  role: 'user'
  parts: GeminiPart[]
}

interface GeminiGenerationConfig {
  responseMimeType: 'application/json'
}

interface GeminiPayload {
  contents: GeminiContent[]
  generationConfig?: GeminiGenerationConfig
}

interface GeminiCandidate {
  content: {
    parts: [{ text: string }]
  }
}

interface GeminiResponse {
  candidates: GeminiCandidate[]
}

export function useGemini() {
  const isLoading = ref<boolean>(false)
  const error = ref<string | null>(null)
  const apiKey = import.meta.env.VITE_GEMINI_API_KEY

  async function callGemini(prompt: string, isJson: boolean = false): Promise<string> {
    isLoading.value = true
    error.value = null

    if (!apiKey) {
      error.value = 'API Key is not configured. Please check your .env file.'
      isLoading.value = false
      throw new Error(error.value)
    }

    const apiUrl = `https://generativelanguage.googleapis.com/v1beta/models/gemini-flash:generateContent?key=${apiKey}`

    // Fix #1: Use the specific GeminiPayload type instead of Record<string, any>
    const payload: GeminiPayload = {
      contents: [{ role: 'user', parts: [{ text: prompt }] }],
    }

    if (isJson) {
      payload.generationConfig = {
        responseMimeType: 'application/json',
      }
    }

    try {
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      })

      if (!response.ok) {
        const errorBody = await response.json()
        throw new Error(
          errorBody.error?.message || `API request failed with status ${response.status}`,
        )
      }

      const result: GeminiResponse = await response.json()

      if (result.candidates?.[0]?.content?.parts?.[0]) {
        return result.candidates[0].content.parts[0].text
      } else {
        throw new Error('Invalid response format from AI.')
      }
    } catch (e: unknown) {
      // Fix #2: Catch error as 'unknown'
      console.error('Gemini API Error:', e)
      // Type guard to safely access the message property
      if (e instanceof Error) {
        error.value = e.message
      } else {
        error.value = 'An unknown error occurred.'
      }
      throw e
    } finally {
      isLoading.value = false
    }
  }

  return {
    isLoading,
    error,
    callGemini,
  }
}
