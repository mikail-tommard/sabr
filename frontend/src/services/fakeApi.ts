export const fakeDelay = <T>(data: T, ms=200) => new Promise<T>(r=>setTimeout(()=>r(structuredClone(data)), ms))
