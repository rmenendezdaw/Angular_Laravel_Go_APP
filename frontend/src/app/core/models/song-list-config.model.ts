export interface SongListConfig {
    type: string;
  
    filters: {
        views?: number,
        release_date?: string
    };
}
  