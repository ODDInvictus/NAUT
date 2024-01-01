import type { Uptime } from '@prisma/client'
import { db } from './db'

export async function getAllApplications() {
  return await db.service.findMany()
}

export async function pingService(slug: string): Promise<Uptime | undefined> {
  const service = await db.service.findFirst({ where: { slug } })

  if (!service) {
    throw new Error(`Service ${slug} onbekend`)
  }

  let uptime = undefined

  if (service.type === 'http') {
    const res = await fetch(service.url)

    uptime = await db.uptime.create({
      data: {
        serviceSlug: service.slug,
        status: res.status === 200
      }
    })

  } else {
    throw new Error(`Ping voor service type ${service.type} nog niet geimplementeerd`)
  }

  return uptime
}