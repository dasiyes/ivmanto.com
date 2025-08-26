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
export async function callGeminiAPI(
  prompt: string,
  generationConfig?: GeminiGenerationConfig,
): Promise<unknown> {
  const apiKey = import.meta.env.VITE_GEMINI_API_KEY
  if (!apiKey) {
    const error =
      'Gemini API key is not configured. Please add VITE_GEMINI_API_KEY to your .env.local file.'
    console.error(error)
    // Return a user-friendly error message from the API call
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
      // Throw a more informative error
      throw new Error(`The AI assistant returned invalid JSON. (Reason: ${errorMessage})`)
    }
  }

  return text
}
