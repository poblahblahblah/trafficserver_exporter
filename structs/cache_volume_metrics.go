// I don't know what to do with these yet and I imagine there will be
// metrics * cache volumes. We don't need it yet, so for now we will
// just ignore them.
type CacheVolumeCounters struct {
    ProxyProcessCacheVolume0BytesUsed                          float64 `json:"proxy.process.cache.volume_0.bytes_used"`
    ProxyProcessCacheVolume0BytesTotal                         float64 `json:"proxy.process.cache.volume_0.bytes_total"`
    ProxyProcessCacheVolume0RAMCacheTotalBytes                 float64 `json:"proxy.process.cache.volume_0.ram_cache.total_bytes"`
    ProxyProcessCacheVolume0RAMCacheBytesUsed                  float64 `json:"proxy.process.cache.volume_0.ram_cache.bytes_used"`
    ProxyProcessCacheVolume0RAMCacheHits                       float64 `json:"proxy.process.cache.volume_0.ram_cache.hits"`
    ProxyProcessCacheVolume0RAMCacheMisses                     float64 `json:"proxy.process.cache.volume_0.ram_cache.misses"`
    ProxyProcessCacheVolume0PreadCount                         float64 `json:"proxy.process.cache.volume_0.pread_count"`
    ProxyProcessCacheVolume0PercentFull                        float64 `json:"proxy.process.cache.volume_0.percent_full"`
    ProxyProcessCacheVolume0LookupActive                       float64 `json:"proxy.process.cache.volume_0.lookup.active"`
    ProxyProcessCacheVolume0LookupSuccess                      float64 `json:"proxy.process.cache.volume_0.lookup.success"`
    ProxyProcessCacheVolume0LookupFailure                      float64 `json:"proxy.process.cache.volume_0.lookup.failure"`
    ProxyProcessCacheVolume0ReadActive                         float64 `json:"proxy.process.cache.volume_0.read.active"`
    ProxyProcessCacheVolume0ReadSuccess                        float64 `json:"proxy.process.cache.volume_0.read.success"`
    ProxyProcessCacheVolume0ReadFailure                        float64 `json:"proxy.process.cache.volume_0.read.failure"`
    ProxyProcessCacheVolume0WriteActive                        float64 `json:"proxy.process.cache.volume_0.write.active"`
    ProxyProcessCacheVolume0WriteSuccess                       float64 `json:"proxy.process.cache.volume_0.write.success"`
    ProxyProcessCacheVolume0WriteFailure                       float64 `json:"proxy.process.cache.volume_0.write.failure"`
    ProxyProcessCacheVolume0WriteBacklogFailure                float64 `json:"proxy.process.cache.volume_0.write.backlog.failure"`
    ProxyProcessCacheVolume0UpdateActive                       float64 `json:"proxy.process.cache.volume_0.update.active"`
    ProxyProcessCacheVolume0UpdateSuccess                      float64 `json:"proxy.process.cache.volume_0.update.success"`
    ProxyProcessCacheVolume0UpdateFailure                      float64 `json:"proxy.process.cache.volume_0.update.failure"`
    ProxyProcessCacheVolume0RemoveActive                       float64 `json:"proxy.process.cache.volume_0.remove.active"`
    ProxyProcessCacheVolume0RemoveSuccess                      float64 `json:"proxy.process.cache.volume_0.remove.success"`
    ProxyProcessCacheVolume0RemoveFailure                      float64 `json:"proxy.process.cache.volume_0.remove.failure"`
    ProxyProcessCacheVolume0EvacuateActive                     float64 `json:"proxy.process.cache.volume_0.evacuate.active"`
    ProxyProcessCacheVolume0EvacuateSuccess                    float64 `json:"proxy.process.cache.volume_0.evacuate.success"`
    ProxyProcessCacheVolume0EvacuateFailure                    float64 `json:"proxy.process.cache.volume_0.evacuate.failure"`
    ProxyProcessCacheVolume0ScanActive                         float64 `json:"proxy.process.cache.volume_0.scan.active"`
    ProxyProcessCacheVolume0ScanSuccess                        float64 `json:"proxy.process.cache.volume_0.scan.success"`
    ProxyProcessCacheVolume0ScanFailure                        float64 `json:"proxy.process.cache.volume_0.scan.failure"`
    ProxyProcessCacheVolume0DirentriesTotal                    float64 `json:"proxy.process.cache.volume_0.direntries.total"`
    ProxyProcessCacheVolume0DirentriesUsed                     float64 `json:"proxy.process.cache.volume_0.direntries.used"`
    ProxyProcessCacheVolume0DirectoryCollision                 float64 `json:"proxy.process.cache.volume_0.directory_collision"`
    ProxyProcessCacheVolume0FragsPerDoc1                       float64 `json:"proxy.process.cache.volume_0.frags_per_doc.1"`
    ProxyProcessCacheVolume0FragsPerDoc2                       float64 `json:"proxy.process.cache.volume_0.frags_per_doc.2"`
    ProxyProcessCacheVolume0FragsPerDoc3                       float64 `json:"proxy.process.cache.volume_0.frags_per_doc.3+"`
    ProxyProcessCacheVolume0ReadBusySuccess                    float64 `json:"proxy.process.cache.volume_0.read_busy.success"`
    ProxyProcessCacheVolume0ReadBusyFailure                    float64 `json:"proxy.process.cache.volume_0.read_busy.failure"`
    ProxyProcessCacheVolume0WriteBytesStat                     float64 `json:"proxy.process.cache.volume_0.write_bytes_stat"`
    ProxyProcessCacheVolume0VectorMarshals                     float64 `json:"proxy.process.cache.volume_0.vector_marshals"`
    ProxyProcessCacheVolume0HdrMarshals                        float64 `json:"proxy.process.cache.volume_0.hdr_marshals"`
    ProxyProcessCacheVolume0HdrMarshalBytes                    float64 `json:"proxy.process.cache.volume_0.hdr_marshal_bytes"`
    ProxyProcessCacheVolume0GcBytesEvacuated                   float64 `json:"proxy.process.cache.volume_0.gc_bytes_evacuated"`
    ProxyProcessCacheVolume0GcFragsEvacuated                   float64 `json:"proxy.process.cache.volume_0.gc_frags_evacuated"`
    ProxyProcessCacheVolume0WrapCount                          float64 `json:"proxy.process.cache.volume_0.wrap_count"`
    ProxyProcessCacheVolume0SyncCount                          float64 `json:"proxy.process.cache.volume_0.sync.count"`
    ProxyProcessCacheVolume0SyncBytes                          float64 `json:"proxy.process.cache.volume_0.sync.bytes"`
    ProxyProcessCacheVolume0SyncTime                           float64 `json:"proxy.process.cache.volume_0.sync.time"`
    ProxyProcessCacheVolume0SpanErrorsRead                     float64 `json:"proxy.process.cache.volume_0.span.errors.read"`
    ProxyProcessCacheVolume0SpanErrorsWrite                    float64 `json:"proxy.process.cache.volume_0.span.errors.write"`
    ProxyProcessCacheVolume0SpanFailing                        float64 `json:"proxy.process.cache.volume_0.span.failing"`
    ProxyProcessCacheVolume0SpanOffline                        float64 `json:"proxy.process.cache.volume_0.span.offline"`
    ProxyProcessCacheVolume0SpanOnline                         float64 `json:"proxy.process.cache.volume_0.span.online"`
}
