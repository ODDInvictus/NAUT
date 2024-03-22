
export function time(time?: Date | null) {
  if (!time) return 'Geen tijd';

  return time.toLocaleTimeString('nl-NL', {
    hour: '2-digit',
    minute: '2-digit',
  });
}

export function date(date?: Date | null) {
  if (!date) return 'Geen datum';

  return date.toLocaleDateString('nl-NL', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  });
}

export function datetime(date?: Date | null) {
  if (!date) return 'Geen datum';

  return date.toLocaleString('nl-NL', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}