import { db } from '$lib/server/db'

export async function getSetting<T>(key: string, defaultValue?: T): Promise<T> {
  const val = await db.settings.findUnique({
    where: {
      key
    }
  })

  if (val === null && defaultValue === undefined) {
    throw new Error(`Setting ${key} not found`)
  }

  if (val === null && defaultValue !== undefined) {
    return defaultValue
  }

  if (val === null) {
    throw new Error(`Setting ${key} not found`)
  }

  return val.value as T
}