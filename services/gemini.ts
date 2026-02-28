// src/services/gemini.ts

// Define a type for the generationConfig if you plan to use it often
interface GeminiGenerationConfig {
  responseMimeType: 'application/json'
  responseSchema: object
}

// Overload for when no `generationConfig` is provided, returning a string.
export async function callGeminiAPI(prompt: string): Promise<string>

// Overload for when a JSON `generationConfig` is provided, returning a generic type T.
export async function callGeminiAPI<T = unknown>(
  prompt: string,
  generationConfig: GeminiGenerationConfig,
): Promise<T>

// Implementation with a combined signature.
// Note: In Nuxt, the API key is passed from the calling component via useRuntimeConfig()
export async function callGeminiAPI(
  prompt: string,
  generationConfig?: GeminiGenerationConfig,
  apiKey?: string,
): Promise<unknown> {
  if (!apiKey) {
    const error =
      'Gemini API key is not configured. Please add NUXT_PUBLIC_GEMINI_API_KEY to your .env file.'
    console.error(error)
    throw new Error('The AI assistant is currently unavailable.')
  }

  const apiUrl = `https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash-preview-05-20:generateContent?key=${apiKey}`

  const payload = {
    contents: [{ role: 'user', parts: [{ text: prompt }] }],
    ...(generationConfig && { generationConfig }),
  }

  const response = await fetch(apiUrl, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  })

  if (!response.ok) {
    const errorBody = await response.text()
    console.error('API Error Body:', errorBody)
    throw new Error(
      `The AI assistant is currently unavailable due to an API error (status: ${response.status}).`,
    )
  }

  const result = await response.json()
  const text = result.candidates?.[0]?.content?.parts?.[0]?.text

  if (!text) {
    throw new Error('Invalid response format from the AI assistant.')
  }

  // If responseMimeType is application/json, parse it before returning
  if (generationConfig?.responseMimeType === 'application/json') {
    try {
      return JSON.parse(text)
    } catch (e: unknown) {
      const errorMessage = e instanceof Error ? e.message : String(e)
      console.error(`Failed to parse JSON response from AI: ${errorMessage}`, text)
      throw new Error(`The AI assistant returned invalid JSON. (Reason: ${errorMessage})`)
    }
  }

  return text
}
