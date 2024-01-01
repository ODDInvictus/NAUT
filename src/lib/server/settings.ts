import { db } from './db';

export async function initSettings() {

}

export async function getSetting(key: string): Promise<string | undefined> {
  const setting = await db.settings.findFirst({
    where: {
      key
    }
  })

  return setting?.value
}