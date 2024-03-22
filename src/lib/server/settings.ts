import { db } from './db';

export async function getSetting(key: string): Promise<string> {
  const setting = await db.settings.findFirst({
    where: {
      key
    }
  })

  if (!setting) {
    throw new Error(`Instelling ${key} niet gevonden`)
  }

  return setting.value
}