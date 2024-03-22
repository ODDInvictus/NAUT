import { getSetting } from '$lib/server/settings';
import type { PageServerLoad } from './$types';

export const load = (async () => {
  const settings = {
    version: await getSetting('VERSION'),
    githubLink: await getSetting('GITHUB_LINK'),
    gitSha: await getSetting('GIT_SHA'),
    buildDate: await getSetting('BUILD_DATE'),
  }

  return settings
}) satisfies PageServerLoad;