import { DEV } from '$env/static/private';
import type { ServerLoadEvent } from '@sveltejs/kit';
import { db } from './db';

type UserData = {
  username: string | null
  name: string | null
  email: string | null
  groups: string[] | undefined
  isDev: boolean | undefined
  isAdmin: boolean | undefined

}

export function parseHeaders(headers: Headers): UserData {
  const groupsString = headers.get('x-authentik-groups')

  const groups = groupsString?.split('|')

  const isAdmin = DEV === 'true'

  return {
    username: headers.get('x-authentik-username'),
    name: headers.get('x-authentik-name'),
    email: headers.get('x-authentik-email'),
    groups: groups,
    isDev: groups?.includes('bakkentrekkers'),
    isAdmin: isAdmin || groups?.includes('ibs-admins')
  }
}

export async function createSession(event: ServerLoadEvent, data: UserData) {

  if (!data.username) {
    return null
  }

  const cookie = event.cookies.get('radio-session')

  console.log(cookie)

  if (cookie) {
    const session = await db.session.findFirst({
      where: {
        id: cookie
      }
    })

    if (!session) {
      return await getSession(cookie)
    }
  }



  // const session = await db.session.findUnique({
  // where: {

  //   }
  // })

  // const user = await db.user.findUnique({
  //   where: {
  //     username: data.username
  //   }
  // })



}

export async function getSession(id: string) {
  return db.session.upsert({
    where: {
      id
    },
    update: {},
    create: {
      id
    }
  })
}